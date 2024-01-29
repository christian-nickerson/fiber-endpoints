from contextlib import asynccontextmanager

import numpy as np
from fastapi import FastAPI, Request, Response
from lightgbm import Booster

from config import settings
from schema import InferenceRequest, InferenceResponse


@asynccontextmanager
async def lifespan(app: FastAPI):
    app.state.model = Booster(model_file=settings.model.file)
    yield


app = FastAPI(lifespan=lifespan)


@app.get("/health")
def health() -> Response:
    """health check endpoint

    :return: Status OK
    """
    return Response("OK")


@app.post("/inference")
def inference(request: Request, body: InferenceRequest) -> InferenceResponse:
    """Handle an inference request

    :param body: Post body
    :return: Prediction response
    """
    inference = np.asarray([body.data])
    prediction = request.app.state.model.predict(inference)
    return InferenceResponse(prediction=prediction)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("app:app", host="0.0.0.0", port=settings.fastapi.port)
