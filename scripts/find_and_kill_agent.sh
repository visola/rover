#!/bin/bash

AGENT_PID=$(ps -le | grep agent | awk '{print $4}')

if [ -z "$AGENT_PID" ]; then
    echo "Agent is not running."
else
    echo "Agent PID: '$AGENT_PID'"
    sudo kill $AGENT_PID
fi

exit 0
