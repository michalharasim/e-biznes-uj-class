from fastapi import FastAPI, Request
from pydantic import BaseModel
import httpx
import uvicorn

app = FastAPI()

LLAMA_API_URL = "http://localhost:11434/api/generate"

class MessageRequest(BaseModel):
    message: str


@app.post("/chat")
async def chat(request: MessageRequest):
    payload = {
        "prompt": request.message,
        "max_tokens": 50,
        "temperature": 0.3,
        "top_p": 0.9,
        "model": "llama3",
        "stream": False
    }

    async with httpx.AsyncClient(timeout=25.0) as client:
        response = await client.post(LLAMA_API_URL, json=payload)
        response.raise_for_status()
        data = response.json()

    answer = data.get("response") or "No response"

    return {"answer": answer}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
