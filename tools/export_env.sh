# copy this command and run it into the main project
# directory to export all variables from .env
export $(grep -v '^#' .env | xargs)