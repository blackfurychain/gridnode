import burn_lock_functions
from burn_lock_functions import GridironchaincliCredentials
from test_utilities import get_required_env_var, get_shell_output


def gridchain_cli_credentials_for_test(key: str) -> GridironchaincliCredentials:
    """Returns GridironchaincliCredentials for the test keyring with from_key set to key"""
    return GridironchaincliCredentials(
        keyring_passphrase="",
        keyring_backend="test",
        from_key=key,
        gridnoded_homedir=f"""{get_required_env_var("HOME")}/.gridnoded"""
    )


def create_new_gridaddr_and_credentials() -> (str, GridironchaincliCredentials):
    new_account_key = get_shell_output("uuidgen")
    credentials = gridchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    credentials.from_key = new_addr["name"]
    return new_addr["address"], credentials,
