import pytest
import gridtool.test_utils


@pytest.fixture(scope="function")
def ctx(request):
    yield from gridtool.test_utils.pytest_ctx_fixture(request)


@pytest.fixture(autouse=True)
def test_wrapper_fixture():
    gridtool.test_utils.pytest_test_wrapper_fixture()
