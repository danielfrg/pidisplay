import os
import sys

import httpx
from typing import Optional
from pydantic import BaseModel
from fastapi import FastAPI, HTTPException

from inkyweb.controller.chrome import take_screenshot

from loguru import logger

logger.remove()
logger.add(
    sys.stdout,
    level="INFO",
    colorize=True,
    format="<green>{time:YYYY-MM-DD HH:mm:ss.SSS}</green> | <level>{level: <8}</level> | {message} | <cyan>{extra}</cyan>",
)

app = FastAPI(title="InkiWeb Controller", version="0.99.0")


DEFAULT_RESOLUTION = (800, 480)
DEFAULT_DISPLAY_URL = os.environ.get("INKYWEB_TARGET", "http://192.168.88.202:8000")


@app.get("/")
async def root():
    return {"message": "InkyWeb Controller API"}


class Update(BaseModel):
    url: str
    target: str | None = None


@app.post("/display/refresh")
async def display_refresh(data: Update):
    image = take_screenshot(data.url, DEFAULT_RESOLUTION)

    async with httpx.AsyncClient() as client:
        target_display = data.target or DEFAULT_DISPLAY_URL
        update_endpoint = f"{target_display}/display/update"

        files = {"image": ("screenshot.png", image, "image/png")}
        response = await client.post(update_endpoint, files=files)

    if response.status_code != 200:
        logger.error("Failed to update display", error=response.json())
        raise HTTPException(
            status_code=response.status_code, detail="Failed to update display API"
        )

    return {"status": "Display updated successfully"}
