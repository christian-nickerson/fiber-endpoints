from dynaconf import Dynaconf

settings = Dynaconf(
    envvar_prefix="FASTAPI",
    settings_files=["settings.toml"],
)
