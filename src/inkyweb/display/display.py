import io
import time
from PIL import Image

from inky.inky_uc8159 import CLEAN


def clear_display(inky):
    for _ in range(2):
        for y in range(inky.height - 1):
            for x in range(inky.width - 1):
                inky.set_pixel(x, y, CLEAN)

        inky.show()
        time.sleep(1.0)


def display_image_from_bytes(inky, image_bytes: bytes, saturation: float = -1):
    image = Image.open(io.BytesIO(image_bytes))

    resized_image = image.resize(inky.resolution)

    if saturation >= 0:
        inky.set_image(resized_image, saturation=saturation)
    else:
        inky.set_image(resized_image)

    inky.show()
