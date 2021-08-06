import json
import requests

host = "127.0.0.1:9000"
with open("docs/swagger.json", 'r') as f:
    temp = json.loads(f.read())
    dict = temp["paths"]
    for k, v in dict.items():
        url = host + k

        data = {}

        method = ""
        if "get" in v.keys():
            method = "GET"
        else:
            method = "POST"
        print(method)

        for k2, v2 in v.items():
            for k3,v3 in v2.items():

                list = v2["parameters"]
                dict3 = list[0]
                print(dict3)
                name = dict3["default"]

                data[name] = name
                print(data)
                resp = {}
                if v2["consumes"] == "multipart/form-data":
                    resp = requests.request(method=method, url=url, data=data)
                elif v2["consumes"] == ".../json":
                    resp = requests.request(method=method, url=url, json=data)
                print(resp)

#
#         for p in v["parameters"]:
#             data[p[name]] = p["name"]
#             print(p)
#             print(data)
#         if v["consumes"] == "multipart/form-data":
#             requests.request(method=method, url=url, data=data)
#         elif v["consumes"] == ".../json":
#             requests.request(method=method, url=url, json=data)
#         resp = requests.request(method=method, url=url, data=data, headers={"Content_Type" : v["consumes"]})
#         print(v["summary"], "status_code is", resp.status_code)
#
#         if v["produces"] == "json":
#             try:
#                 resp.json()
#             except:
#                 print(k, "response is error")