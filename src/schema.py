from typing import List

from pydantic import BaseModel


class InferenceRequest(BaseModel):
    data: List[float]


class InferenceResponse(BaseModel):
    prediction: float
