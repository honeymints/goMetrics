import aiohttp
import asyncio
#


async def parse(city_id: int, req_type: str):
    async with aiohttp.ClientSession() as session:
        link = 'http://opendata.kz/api/sensor/getListWithLastHistory?cityId={city_id}'.format(city_id=city_id)
        async with session.get(link) as resp:
            res_json = await resp.json()

    metrics_dict = {}
    for sensor in res_json['sensors']:
        if req_type == "temperature":
            metrics_dict.update({
                sensor['name']: {
                    'temperature': sensor['history'][0]['data']['field3'],
                    'humidity': sensor['history'][0]['data']['field5']
                }})
        else:
            metrics_dict.update({
                sensor['name']: {
                    'co2': sensor['history'][0]['data']['field1'],
                    'pm25': sensor['history'][0]['data']['field2']
                }})

    return metrics_dict


    # def update_city_lists(city_lists, metrics_dict):
    #     for city_list in city_lists:
    #         for i, district in enumerate(city_list):
    #             updated_district = {key: metrics_dict[key] for key in district if key in metrics_dict}
    #             city_list[i] = updated_district
    #     return city_lists
    #
    # # usage
    # all_city_lists = [astana_list, oskemen_list, atyrau_list, semei_list]
    # updated_city_lists = update_city_lists(all_city_lists, metrics_dict)
    # print(updated_city_lists)

    # city_districts = city_dict.get(city_id)
    #
    # for district in city_districts:
    #     sensor_name = district["name"]
    #     if sensor_name in metrics_dict:
    #         district.update(metrics_dict[sensor_name])

    # print(city_districts)
    # for i in city_districts:
    #     for j in metrics_dict:
    #         if j in i:
    #             city_districts[i] = metrics_dict[j]

    # for sensor in metrics_dict:
    #     for district in city_districts:
    #         if sensor in district:


# astana_esil = {"NUR02", "NUR04", "NUR05", "NUR11", "NUR14", "NUR20", "NUR22","NUR24","NUR27","NUR28","NUR31","NUR32"
#                "NUR33", "NUR34","NUR41","NUR45","NUR49","NURTEST001"}
# astana_saryarka = {"NUR10","NUR21","NUR35","NUR40","NUR43","NUR44"}
# astana_baikonur = {"NUR08", "NUR09", "NUR12","NUR13", "NUR15","NUR16","NUR18","NUR19","NUR23","NUR26","NUR30"}
# astana_almatynskiy = {"NUR01", "NUR03", "NUR07","NUR25","NUR29","NUR36","NUR37","NUR38","NUR39","NUR46","NUR48",
#                       "NUR50"}
#
# astana_list = [{"name": name} for name in astana_almatynskiy.union(astana_baikonur, astana_esil, astana_saryarka)]
#
# oskemen_zashita = {"OSK002"}
# oskemen_ulbinsk = {"OSK001", "OSK003","OSK005"}
# oskemen_zavodsk = {"OSK008"}
# oskemen_tsentr = {"OSK006", "OSK009"}
# oskemen_ksht = {"OSK004"}
#
#
# oskemen_list = [{"name": name} for name in oskemen_ksht.union(oskemen_tsentr,oskemen_zavodsk,oskemen_ulbinsk,oskemen_zashita)]
#
# atyrau_zapad = {"ATR002 (БС-2)", "ATR004 (БС-22)", "ATR007 (БС-12)", "ATR009 (БС-25)"}
# atyrau_vostok = {"ATR001 (БС-1)","ATR003 (БС-16)","ATR005 (БС-5)", "ATR006 (БС-33)", "ATR008 (БС-8)",
#                  "ATR010 (БС-27)"}
#
# atyrau_list = [{"name": name} for name in atyrau_vostok.union(atyrau_zapad)]
#
# semei_sever = {"SEM06 (БС10)","SEM07 (БС11)", "SEM08 (БС12)","SEM08 (БС3)", "SEM09 (БС13)", "SEM10 (БС15)",
#                "SEM11 (ЗИП)", "SEM12 (ЗИП)", "SEM13 (ЗИП)","SEM14 (ТЕСТ)"}
# semei_yug = {"SEM01 (БС1)", "SEM02 (БС4)", "SEM03 (БС5)", "SEM04 (БС7)", "SEM05 (БС9)"}
#
# semei_list = [{"name": name} for name in semei_yug.union(semei_sever)]
#
# city_dict = {
#     1: astana_list,
#     2: oskemen_list,
#     3: atyrau_list,
#     4: semei_list
# }
# async def main():
#     await parse(1, "temperature")


# loop = asyncio.get_event_loop()
# loop.run_until_complete(main())
