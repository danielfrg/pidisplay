import os

from pydantic import BaseModel
from fastapi import FastAPI, UploadFile, File, HTTPException

from loguru import logger


app = FastAPI(title="InkiWeb Display", version="0.99.0")


DEFAULT_DISPLAY_URL = os.environ.get("INKYWEB_TARGET", "http://localhost")


@app.get("/")
async def root():
    return {"message": "InkyWeb Display API"}


# class Update(BaseModel):
#     url: str
#     target: Optional[str]


@app.post("/display/update")
async def update_display(image: UploadFile = File(...)):
    """
    Receives an image file and updates the eInk display.
    """

    try:
        image_bytes = await image.read()
    except Exception as e:
        logger.error("Error reading uploaded image: %s", e)
        raise HTTPException(status_code=400, detail="Invalid image data")

    if not update_display(image_bytes):
        raise HTTPException(status_code=500, detail="Display update failed")

    return {"message": "Display updated successfully"}
