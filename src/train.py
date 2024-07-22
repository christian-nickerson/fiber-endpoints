from lightgbm import LGBMRegressor

from config import settings
from data.sklearn import generate_data


def train() -> None:
    """Train LGBM Model and save to file"""
    X, y = generate_data(samples=settings.model.training_samples)
    model = LGBMRegressor(n_estimators=settings.model.estimators)
    model.fit(X, y)
    model.booster_.save_model(settings.model.file, num_iteration=model.best_iteration_)


if __name__ == "__main__":
    train()
