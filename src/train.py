from lightgbm import LGBMRegressor
from sklearn.datasets import make_regression

from config import settings


def train() -> None:
    """Train LGBM Model and save to file"""
    X, y = make_regression(
        n_samples=100_000, n_features=20, n_informative=15, noise=0.1, random_state=5
    )
    model = LGBMRegressor(n_estimators=500)
    model.fit(X, y)
    model.booster_.save_model(settings.model.file, num_iteration=model.best_iteration_)


if __name__ == "__main__":
    train()
