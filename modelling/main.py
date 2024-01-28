from sklearn.datasets import make_regression
from lightgbm import LGBMRegressor

if __name__ == "__main__":

    X, y = make_regression(n_samples=100_000, n_features=20, n_informative=15, noise=0.1, random_state=5)
    model = LGBMRegressor(n_estimators=500)
    model.fit(X, y)
    model.booster_.save_model("model.txt", num_iteration=model.best_iteration_)