from numpy import ndarray
from sklearn.datasets import make_regression


def generate_data(samples: int) -> tuple[ndarray, ndarray, ndarray]:
    """generate regression data"""
    return make_regression(
        n_samples=samples, n_features=20, n_informative=15, noise=0.1, random_state=5
    )
