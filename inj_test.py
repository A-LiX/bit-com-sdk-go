"""
基于 inject python sdk，他们提供 grpc style 的调用

https://github.com/InjectiveLabs/sdk-python
"""
#!/usr/bin/python
import asyncio
import datetime
import time
import csv
import os
from pyinjective.composer import Composer as ProtoMsgComposer
from pyinjective.async_client import AsyncClient
from pyinjective.transaction import Transaction
from pyinjective.constant import Network
from pyinjective.wallet import PrivateKey
from typing import Dict, Tuple, List, NamedTuple

# import pyinjective.proto.exchange.injective_accounts_rpc_pb2 as accounts_rpc_pb
# import pyinjective.proto.exchange.injective_accounts_rpc_pb2_grpc as accounts_rpc_grpc

LEVERAGE = 3


class OrderInfoParam(NamedTuple):
    order_id: str
    client_order_id: str or None
    symbol: str


class TradeLimitParam(NamedTuple):
    price: float
    amount: float
    symbol: str
    client_order_id: str or None
    side: str


MARKET_INFO_DICT = {'injective': {
    'WBTC': {'id': '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599',
             'denom': 'peggy0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599', 'decimals': 8},
    'GF': {'id': '0xAaEf88cEa01475125522e117BFe45cF32044E238',
           'denom': 'peggy0xAaEf88cEa01475125522e117BFe45cF32044E238', 'decimals': 18},
    'QNT': {'id': '0x4a220E6096B25EADb88358cb44068A3248254675',
            'denom': 'peggy0x4a220E6096B25EADb88358cb44068A3248254675', 'decimals': 18},
    'WETH': {'id': '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2',
             'denom': 'peggy0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', 'decimals': 18},
    'USDT': {'id': '0xdAC17F958D2ee523a2206206994597C13D831ec7',
             'denom': 'peggy0xdAC17F958D2ee523a2206206994597C13D831ec7', 'decimals': 6},
    'MATIC': {'id': '0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0',
              'denom': 'peggy0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0', 'decimals': 18},
    'LINK': {'id': '0x514910771AF9Ca656af840dff83E8264EcF986CA',
             'denom': 'peggy0x514910771AF9Ca656af840dff83E8264EcF986CA', 'decimals': 18},
    'AXS': {'id': '0xBB0E17EF65F82Ab018d8EDd776e8DD940327B28b',
            'denom': 'peggy0xBB0E17EF65F82Ab018d8EDd776e8DD940327B28b', 'decimals': 18},
    'GRT': {'id': '0xc944E90C64B2c07662A292be6244BDf05Cda44a7',
            'denom': 'peggy0xc944E90C64B2c07662A292be6244BDf05Cda44a7', 'decimals': 18},
    'AAVE': {'id': '0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9',
             'denom': 'peggy0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9', 'decimals': 18},
    'UST': {'id': '0xa47c8bf37f92aBed4A126BDA807A7b7498661acD',
            'denom': 'ibc/B448C0CA358B958301D328CCDC5D5AD642FC30A6D3AE106FF721DB315F3DDE5C', 'decimals': 6},
    'ATOM': {'id': '0x8D983cb9388EaC77af0474fA441C4815500Cb7BB',
             'denom': 'ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9', 'decimals': 6},
    'LUNA': {'id': '0xd2877702675e6cEb975b4A1dFf9fb7BAF4C91ea9',
             'denom': 'ibc/B8AF5D92165F35AB31F3FC7C7B444B9D240760FA5D406C49D24862BD0284E395', 'decimals': 6},
    'INJ': {'id': '0xe28b3B32B6c345A34Ff64674606124Dd5Aceca30', 'denom': 'inj', 'decimals': 18},
    'SUSHI': {'id': '0x6B3595068778DD592e39A122f4f5a5cF09C90fE2',
              'denom': 'peggy0x6B3595068778DD592e39A122f4f5a5cF09C90fE2', 'decimals': 18},
    'UNI': {'id': '0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984',
            'denom': 'peggy0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984', 'decimals': 18},
    'USDC': {'id': '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
             'denom': 'peggy0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', 'decimals': 6},
    'SNX': {'id': '0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F',
            'denom': 'peggy0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F', 'decimals': 18},
    'WETH_USDC': {'id': '0x01e920e081b6f3b2e5183399d5b6733bb6f80319e6be3805b95cb7236910ff0e'},
    'INJ_USDC': {'id': '0xe0dc13205fb8b23111d8555a6402681965223135d368eeeb964681f9ff12eb2a'},
    'USDT_USDC': {'id': '0x8b1a4d3e8f6b559e30e40922ee3662dd78edf7042330d4d620d188699d1a9715'},
    'LINK_USDC': {'id': '0xfe93c19c0a072c8dd208b96694e024305a7dff01bbf12cac2bfa81b246c69040'},
    'AAVE_USDC': {'id': '0xcdfbfaf1f24055e89b3c7cc763b8cb46ffff08cdc38c999d01f58d64af75dca9'},
    'MATIC_USDC': {'id': '0x5abfffe9079d53e0bf8ee9b3064b427acc3d71d6ba58a44235abe38f60115678'},
    'UNI_USDT': {'id': '0xe8bf0467208c24209c1cf0fd64833fa43eb6e8035869f9d043dbff815ab76d01'},
    'SUSHI_USDC': {'id': '0x9a629b947b6f946af4f6076cfda67f3535d73ee3cef6176cf6d9c8d6b0a03f37'},
    'GRT_USDC': {'id': '0xa43d2be9861efb0d188b136cef0ae2150f80e08ec318392df654520dd359fcd7'},
    'INJ_USDT': {'id': '0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0'},
    'MATIC_USDT': {'id': '0x28f3c9897e23750bf653889224f93390c467b83c86d736af79431958fff833d1'},
    'UNI_USDC': {'id': '0x09cc2c28fbedbdd677e07924653f8f583d0ee5886e74046e7f114210d990784b'},
    'LINK_USDT': {'id': '0x26413a70c9b78a495023e5ab8003c9cf963ef963f6755f8b57255feb5744bf31'},
    'WETH_USDT': {'id': '0xd1956e20d74eeb1febe31cd37060781ff1cb266f49e0512b446a5fafa9a16034'},
    'AAVE_USDT': {'id': '0x01edfab47f124748dc89998eb33144af734484ba07099014594321729a0ca16b'},
    'GRT_USDT': {'id': '0x29255e99290ff967bc8b351ce5b1cb08bc76a9a9d012133fb242bdf92cd28d89'},
    'SUSHI_USDT': {'id': '0x0c9f98c99b23e89dbf6a60bec05372790b39e03da0f86dd0208fc8e28751bd8c'},
    'SNX_USDT': {'id': '0x51092ddec80dfd0d41fee1a7d93c8465de47cd33966c8af8ee66c14fe341a545'},
    'QNT_USDT': {'id': '0xbe9d4a0a768c7e8efb6740be76af955928f93c247e0b3a1a106184c6cf3216a7'},
    'WBTC_USDC': {'id': '0x170a06eb653548f67e94b0fcb82c5258c83b0a2b62ed24c55749d5ac77bc7621'},
    'AXS_USDT': {'id': '0x7471d361b90fc8541267bd088f498c2a461a2c0c57ff2b9a08279480e803b470'},
    'ATOM_USDT': {'id': '0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa'},
    'GF_USDT': {'id': '0x7f71c4fba375c964be8db7fc7a5275d974f8c6cdc4d758f2ac4997f106bb052b'},
    'UST_USDT': {'id': '0x0f1a11df46d748c2b20681273d9528021522c6a0db00de4684503bbd53bef16e'},
    'LUNA_UST': {'id': '0xdce84d5e9c4560b549256f34583fb4ed07c82026987451d5da361e6e238287b3'},
    'INJ_UST': {'id': '0xfbc729e93b05b4c48916c1433c9f9c2ddb24605a73483303ea0f87a8886b52af'}
},
    'injectiveperpetual': {
        'BTC_USDT_SWAP': {'id': '0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce'},
        'ETH_USDT_SWAP': {'id': '0x54d4505adef6a5cef26bc403a33d595620ded4e15b9e2bc3dd489b714813366a'},
        'BNB_USDT_SWAP': {'id': '0x1c79dac019f73e4060494ab1b4fcba734350656d6fc4d474f6a238c13c6f9ced'},
        'INJ_USDT_SWAP': {'id': '0x9b9980167ecc3645ff1a5517886652d94a0825e54a77d2057cbbe3ebee015963'},
        'LUNA_UST_SWAP': {'id': '0x8158e603fb80c4e417696b0e98765b4ca89dcf886d3b9b2b90dc15bfb1aebd51'},
        'ATOM_USDT_SWAP': {'id': '0xc559df216747fc11540e638646c384ad977617d6d8f0ea5ffdfc18d52e58ab01'}
    }}


class InjectiveperpetualApi():
    """
    api 初始化后需要调用 init 获取初试账号 sequence，后续发消息需要使用
    >> api = InjectiveApi(...)
    >> await api.init()

    初始化的参数：
    - pub_key       不填，inj 不需要
    - secret_key    钱包的私钥
    - fee_recipient 返还费用的地址

    实现的方法有：
    - buy_limit
    - sell_limit
    - get_orders
    - cancel_order
    - batch_cancel
    - get_position

    具体调用方法和返回实例可以参考同目录下 test_async_api 文件，里面有每个方法的单元测试
    """

    # SEND_LOCK = asyncio.Lock()
    # REQ_NUM_MAX = 1

    def __init__(self, pub_key, secret_key, fee_recipient):
        self.SEND_LOCK = asyncio.Lock()
        self.network = Network.mainnet(node='sentry3')
        self.composer = ProtoMsgComposer(network=self.network.string())
        self.client = AsyncClient(network=self.network, insecure=True)
        self.private_key = PrivateKey.from_hex(secret_key)
        self.public_key = self.private_key.to_public_key()
        self.address = self.public_key.to_address()
        self.subaccount_id = self.address.get_subaccount_id(index=0)
        self.fee_recipient = fee_recipient


    async def _close_node(self):
        await self.client.chain_channel.close()
        await self.client.exchange_channel.close()

    async def change_sentry_node(self, node):
        nodes = [
            'sentry0',  # us, prod
            'sentry1',  # us, prod
            'sentry2',  # us, staging
            'sentry3',  # tokyo, prod,
            'sentry4',
            'sentry.cd',  # dedicated github-runner
            'blind',
            'asymm_inner_node',
            'asymm_outer_node',
        ]
        if node not in nodes:
            raise ValueError("must be one of {}".format(nodes))

        await self.client.chain_channel.close()
        await self.client.exchange_channel.close()

        inj_network = Network.mainnet(node=node)
        inj_client = AsyncClient(network=inj_network, insecure=True)
        print('sentry change to: ', node)
        self.network = inj_network
        self.client = inj_client
        self.composer = ProtoMsgComposer(network=self.network.string())
        return

    async def batch_update(self,
                           spot_orders_to_cancel: List[OrderInfoParam] = None,
                           derivative_orders_to_cancel: List[OrderInfoParam] = None,
                           spot_orders_to_create: List[TradeLimitParam] = None,
                           derivative_orders_to_create: List[TradeLimitParam] = None,
                           spot_market_ids_to_cancel_all: List[OrderInfoParam] = None,
                           derivative_market_ids_to_cancel_all: List[OrderInfoParam] = None,
                           ):

        # market_info改成 spot 和 deribative 混合的
        created_res = []
        spot_orders_to_cancel_orders = [
            self.composer.OrderData(
                market_id=MARKET_INFO_DICT['injective'].get(order_param.symbol).get('id'),
                subaccount_id=self.subaccount_id,
                order_hash=order_param.order_id
            )
            for order_param in spot_orders_to_cancel
        ] if spot_orders_to_cancel else []
        derivative_orders_to_cancel_orders = [
            self.composer.OrderData(
                market_id=MARKET_INFO_DICT['injectiveperpetual'].get(order_param.symbol).get('id'),
                subaccount_id=self.subaccount_id,
                order_hash=order_param.order_id
            )
            for order_param in derivative_orders_to_cancel
        ] if derivative_orders_to_cancel else []
        spot_orders_to_create_orders = [
            self.composer.SpotOrder(
                market_id=MARKET_INFO_DICT['injective'].get(order_param.symbol).get('id'),
                subaccount_id=self.subaccount_id,
                fee_recipient=self.fee_recipient,
                price=float(order_param.price),
                quantity=float(order_param.amount),
                is_buy=True if order_param.side == 'buy' else False
            )
            for order_param in spot_orders_to_create
        ] if spot_orders_to_create else []
        derivative_orders_to_create_orders = [
            self.composer.DerivativeOrder(
                market_id=MARKET_INFO_DICT['injectiveperpetual'].get(order_param.symbol).get('id'),
                subaccount_id=self.subaccount_id,
                fee_recipient=self.fee_recipient,
                price=float(order_param.price),
                quantity=float(order_param.amount),
                leverage=LEVERAGE,
                is_buy=True if order_param.side == 'buy' else False,
            )
            for order_param in derivative_orders_to_create
        ] if derivative_orders_to_create else []

        spot_market_ids_to_cancel_all_orders = {
            MARKET_INFO_DICT['injective'].get(order_param.symbol).get('id')
            for order_param in spot_market_ids_to_cancel_all
        } if spot_market_ids_to_cancel_all else []

        derivative_market_ids_to_cancel_all_orders = {
            MARKET_INFO_DICT['injectiveperpetual'].get(order_param.symbol).get('id')
            for order_param in derivative_market_ids_to_cancel_all
        } if derivative_market_ids_to_cancel_all else []
        msg = self.composer.MsgBatchUpdateOrders(
            sender=self.address.to_acc_bech32(),
            subaccount_id=self.subaccount_id if spot_market_ids_to_cancel_all_orders or derivative_market_ids_to_cancel_all_orders else None,
            derivative_orders_to_create=derivative_orders_to_create_orders,
            spot_orders_to_create=spot_orders_to_create_orders,
            derivative_orders_to_cancel=derivative_orders_to_cancel_orders,
            spot_orders_to_cancel=spot_orders_to_cancel_orders,
            spot_market_ids_to_cancel_all=list(spot_market_ids_to_cancel_all_orders),
            derivative_market_ids_to_cancel_all=list(derivative_market_ids_to_cancel_all_orders),
        )
        sim_resp = await self._send_msg(msg)
        for i in sim_resp.derivative_order_hashes:
            single_res = {
                "derivative_orderId": i,
                "client_order_id": None,
                "symbol": None,
            }
            created_res.append(single_res)
        for i in sim_resp.spot_order_hashes:
            single_res = {
                "spot_orderId": i,
                "client_order_id": None,
                "symbol": None,
            }
            created_res.append(single_res)
        spot_cancel_success_list = sim_resp.spot_cancel_success
        derivative_cancel_success_list = sim_resp.derivative_cancel_success

        spot_cancel_id_list = [i.order_id for i in
                               spot_orders_to_cancel] if spot_orders_to_cancel else []
        derivative_cancel_id_list = [i.order_id for i in
                                     derivative_orders_to_cancel] if derivative_orders_to_cancel else []
        print(created_res, spot_cancel_success_list, derivative_cancel_success_list, spot_cancel_id_list, derivative_cancel_id_list)
        return created_res, spot_cancel_success_list, derivative_cancel_success_list, spot_cancel_id_list, derivative_cancel_id_list

    async def _send_msg(self, msg):

        # async with self.SEND_LOCK:
        if not self.address.sequence:
            await self.address.async_init_num_seq(self.network.lcd_endpoint)

        tx = (
            Transaction()
                .with_messages(msg)
                .with_sequence(self.address.get_sequence())
                .with_account_num(self.address.get_number())
                .with_chain_id(self.network.chain_id)
            # .with_timeout_height(current_height + 50)
        )

        sim_sign_doc = tx.get_sign_doc(self.public_key)
        sim_sig = self.private_key.sign(sim_sign_doc.SerializeToString())
        sim_tx_raw_bytes = tx.get_tx_data(sim_sig, self.public_key)

        # simulate tx
        (simRes, success) = await self.client.simulate_tx(sim_tx_raw_bytes)
        # sim_res_msg = ProtoMsgComposer.MsgResponses(simRes.result.data, simulation=True)

        if not success:
            if 'incorrect account sequenc' in str(simRes):
                self.address.sequence = self._parse_sequence(str(simRes))

            raise Exception('simulation error' + str(simRes))

        # build tx
        gas_price = 500000000
        gas_limit = simRes.gas_info.gas_used + 25000  # add 15k for gas, fee computation
        fee = [self.composer.Coin(
            amount=str(gas_price * gas_limit),
            denom=self.network.fee_denom,
        )]
        tx = tx.with_gas(gas_limit).with_fee(fee).with_memo("").with_timeout_height(0)
        sign_doc = tx.get_sign_doc(self.public_key)
        sig = self.private_key.sign(sign_doc.SerializeToString())
        tx_raw_bytes = tx.get_tx_data(sig, self.public_key)

        # broadcast tx: send_tx_async_mode, send_tx_sync_mode, send_tx_block_mode
        res = await self.client.send_tx_block_mode(tx_raw_bytes)
        if res.code != 0:
            if 'incorrect account sequenc' in str(res):
                self.address.sequence = self._parse_sequence(str(res))
            raise Exception('broadcast error' + str(res.raw_log))

        sim_res_msg = ProtoMsgComposer.MsgResponses(simRes.result.data, simulation=True)
        if sim_res_msg:
            return sim_res_msg[0]  # sell/buy order return hash
        else:
            return True  # candel order return True

    def _parse_sequence(self, arg):
        rslt = int(arg[arg.find('expected ') + 9: arg.find(", got")])
        return rslt


if __name__ == '__main__':

    file = open('./log.csv','w', encoding='utf-8')
    csv_writer = csv.writer(file)
    for i in range(1,1001):
        inj =InjectiveperpetualApi(
            
            fee_recipient='inj18ul3vv8lk6zkv34wx4pja27sdmguewhll622m6')
        batch_update_func = inj.batch_update(
            # spot_orders_to_cancel=[OrderInfoParam('0x3e166ee0adfa72745de75a34a386d898c890efaea7e4798e911400a752a085c0',None,'INJ_USDT')],
            # derivative_orders_to_cancel=[
            #     OrderInfoParam('0x03423821063f997df5c74f3a842fb0178ac595a1d41aa1209a7f064f789da9f9',None,'BTC_USDT_SWAP'),
            #     OrderInfoParam('0xfb114412a2746fba37c812326eedc1c6818ae2a89fb08896aab8ee8b3d0ba451',None,'BTC_USDT_SWAP'),
            # ],
            spot_orders_to_create=[TradeLimitParam(3.0, 0.02, "INJ_USDT", client_order_id=None, side='buy')],
            derivative_orders_to_create=[
                TradeLimitParam(49500, 0.0001, "BTC_USDT_SWAP", client_order_id=None, side='sell'),
                TradeLimitParam(49600, 0.0001, "BTC_USDT_SWAP", client_order_id=None, side='sell'),
            ],
            spot_market_ids_to_cancel_all=[OrderInfoParam('0x95c7c6459a00e6683f55119987fb02a92e5905a1182f533944a47f4bba9368e1',None,'INJ_USDT')],
            derivative_market_ids_to_cancel_all=[OrderInfoParam('0x9569b6a05aa09c44da8947618a2d46fba1063c83d837133d6ff26f9298bcccdd',None,'BTC_USDT_SWAP')],
        )
        loop = asyncio.get_event_loop() or asyncio.new_event_loop()
        # loop = asyncio.new_event_loop()
 
        start_time = datetime.datetime.now()
        loop.run_until_complete(batch_update_func)
        end_time = datetime.datetime.now()
        delta = end_time - start_time
        delta_gmtime = time.gmtime(delta.total_seconds())
        duration_str = time.strftime("%S", delta_gmtime)
        print ("duration:", duration_str)
        csv_writer.writerow([duration_str])
        #i += 1
    file.close()







