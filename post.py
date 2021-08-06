import httpx

data = {"name": "name"}
text = httpx.post("http://127.0.0.1:9000/v1/islike", data=data).json()
print(text)


