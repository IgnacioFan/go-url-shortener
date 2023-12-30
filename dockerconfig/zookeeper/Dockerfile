FROM zookeeper:3.8.2

COPY . .

# 2181 defualt client port
# 2888 default peer port
# 3888 default leader election port
# 9876 admin server port
EXPOSE 2181 2888 3888 9876

# Start Zookeeper when the container starts
CMD ["bin/zkServer.sh", "start-foreground"]
