from contextlib import asynccontextmanager

import numpy as np
from fastapi import FastAPI, Request, Response
from lightgbm import Booster
from opentelemetry import trace

from config import settings
from otel import instrumentor, oltp_tracer
from schema import InferenceRequest, InferenceResponse


@asynccontextmanager
async def lifespan(app: FastAPI):
    app.state.model = Booster(model_file=settings.model.file)
    yield


app = FastAPI(lifespan=lifespan)
instrumentor.instrument_app(app, tracer_provider=oltp_tracer)
tracer = trace.get_tracer(__name__)


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
    with tracer.start_as_current_span("model-inference"):
        prediction = request.app.state.model.predict(inference)
    return InferenceResponse(prediction=prediction)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run("app:app", host="0.0.0.0", port=settings.fastapi.port)
