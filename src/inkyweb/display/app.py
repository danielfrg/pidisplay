import os
import sys

from pydantic import BaseModel
from fastapi import FastAPI, UploadFile, File, HTTPException
from inky.auto import auto

from inkyweb.display.display import display_image_from_bytes

from loguru import logger

logger.remove()
logger.add(
    sys.stdout,
    level="INFO",
    colorize=True,
    format="<green>{time:YYYY-MM-DD HH:mm:ss.SSS}</green> | <level>{level: <8}</level> | {message} | <cyan>{extra}</cyan>",
)


DEFAULT_DISPLAY_URL = os.environ.get("INKYWEB_TARGET", "http://localhost")


app = FastAPI(title="InkiWeb Display", version="0.99.0")


# inky = auto(ask_user=True, verbose=True)
inky = auto()
print("Display resolution:", inky.resolution)


@app.get("/")
async def root():
    return {"message": "InkyWeb Display API"}


# class Update(BaseModel):
#     url: str
#     target: Optional[str]


@app.post("/display/update")
async def update_display(image: UploadFile = File(...)):
    """
    Receives an image file and updates the eInk display
    """

    try:
        image_bytes = await image.read()
    except Exception as e:
        logger.error("Error reading uploaded image", error=e)
        raise HTTPException(status_code=400, detail="Invalid image data")

    if not display_image_from_bytes(inky, image_bytes):
        raise HTTPException(status_code=500, detail="Display update failed")

    return {"message": "Display updated"}
