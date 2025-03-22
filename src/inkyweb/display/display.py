import io
from PIL import Image


def display_image_from_bytes(inky, image_bytes: bytes, saturation: float = -1):
    image = Image.open(io.BytesIO(image_bytes))

    resized_image = image.resize(inky.resolution)

    if saturation >= 0:
        inky.set_image(resized_image, saturation=saturation)
    else:
        inky.set_image(resized_image)

    inky.show()
