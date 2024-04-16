#!/bin/bash

# Step 1: Set the maximum file limit globally
echo "fs.file-max = 999999" | sudo tee -a /etc/sysctl.conf

# Step 2: Reload the sysctl configuration
sudo sysctl -p

# Step 3: Run httpflood with the desired parameters
sudo ./httpflood "$@"
