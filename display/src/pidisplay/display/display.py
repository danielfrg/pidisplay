import io
import time
from PIL import Image, ImageOps

from inky.inky_uc8159 import CLEAN


def clear_display(inky):
    for y in range(inky.height - 1):
        for x in range(inky.width - 1):
            inky.set_pixel(x, y, CLEAN)

    inky.show()


def display_image_from_bytes(inky, image_bytes: bytes, saturation: float = -1, rotation: int = 0):
    image = Image.open(io.BytesIO(image_bytes))

    # Apply rotation if specified
    if rotation == 90:
        image = image.transpose(Image.ROTATE_90)
    elif rotation == 180:
        image = image.transpose(Image.ROTATE_180)
    elif rotation == 270:
        image = image.transpose(Image.ROTATE_270)

    resized_image = image.resize(inky.resolution)

    if saturation >= 0:
        inky.set_image(resized_image, saturation=saturation)
    else:
        inky.set_image(resized_image)

    inky.show()
