# PREREQUISITES
#
# export GOPATH=$HOME/mygo
# export PATH=$PATH:$GOPATH/bin
#
# Above variables should be set in order to run the
# deployment script.
# Normally, I put them in ~/.bashrc
#
# Configures the PROJECT_SOURCE and
# EXECUTABLE variables 
# to reflect stuff on the server

PROJECT_SOURCE=$GOPATH/src/github.com/dst-hackathon/socialradar-api
EXECUTABLE=socialradar-api

echo -e "\n"
PID_FILE="$GOPATH/API_ID.tmp"
API_ID=`cat $PID_FILE`
echo "Current PID: $API_ID"

cd $PROJECT_SOURCE

echo "Terminating current process..."
kill $API_ID
echo -e "Current process terminated.\n"

echo "Pulling code from repository..."
git pull
echo -e "Pull finished.\n"

echo "Installing..."
go install
echo -e "Go install done.\n"

echo "Running process.."
$EXECUTABLE &
echo -e "Process started.\n"

API_ID=$!
echo $API_ID > $PID_FILE
echo "Current PID: $API_ID"

echo -e "Application deployed.\n"