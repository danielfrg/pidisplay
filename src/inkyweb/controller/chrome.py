import os
import tempfile
import subprocess

from typing import Optional, Tuple

from loguru import logger


def take_screenshot(
    url_or_path: str, dim: Tuple[int, int], binary="chrome-headless-shell"
) -> Optional[bytes]:
    with tempfile.NamedTemporaryFile(suffix=".png", delete=False) as img_file:
        img_file_path = img_file.name

    logger.debug("Taking screenshot of %s to %s", url=url_or_path, path=img_file_path)

    command = [
        "chrome-headless-shell",
    ]

    if binary == "chromium-browser":
        command = [
            "chromium-browser",
            "--headless=old",
        ]

    command.extend(
        [
            f"--screenshot={img_file_path}",
            f"--window-size={dim[0]},{dim[1]}",
            "--no-sandbox",
            "--disable-gpu",
            "--disable-software-rasterizer",
            "--disable-dev-shm-usage",
            "--hide-scrollbars",
            url_or_path,
        ]
    )

    result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    if result.returncode != 0 or not os.path.exists(img_file_path):
        logger.error("Failed to take screenshot:", result.stderr.decode("utf-8"))
        return None

    try:
        with open(img_file_path, "rb") as f:
            image_bytes = f.read()
    except Exception as e:
        logger.error("Error reading screenshot file: %s", e)
        return None
    finally:
        try:
            os.remove(img_file_path)
        except Exception as e:
            logger.warning("Error cleaning up screenshot file: %s", e)

    return image_bytes
