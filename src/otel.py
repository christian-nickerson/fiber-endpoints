from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

from config import settings

resource = Resource.create({SERVICE_NAME: settings.fastapi.name})
provider = TracerProvider(resource=resource)
batch_span = BatchSpanProcessor(OTLPSpanExporter(f"http://{settings.otel.host}"))

trace.set_tracer_provider(provider)
oltp_tracer = trace.get_tracer_provider().add_span_processor(batch_span)

instrumentor = FastAPIInstrumentor()
