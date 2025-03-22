import os

import httpx
from typing import Optional
from pydantic import BaseModel
from fastapi import FastAPI, HTTPException


from .chrome import take_screenshot


app = FastAPI(title="InkiWeb Controller", version="0.99.0")


DEFAULT_DISPLAY_URL = os.environ.get("INKYWEB_TARGET", "http://192.168.88.48")


@app.get("/")
async def root():
    return {"message": "InkyWeb Controller API"}


class Update(BaseModel):
    url: str
    target: Optional[str]


@app.post("/display/refresh")
async def display_refresh(data: Update):
    image = take_screenshot(data.url, (800, 600))

    async with httpx.AsyncClient() as client:
        files = {"file": ("screenshot.png", image, "image/png")}
        response = await client.post(DEFAULT_DISPLAY_URL, files=files)

    if response.status_code != 200:
        raise HTTPException(
            status_code=response.status_code, detail="Failed to update display API"
        )

    return {"status": "Display updated successfully"}
