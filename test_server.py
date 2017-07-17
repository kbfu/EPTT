# coding=utf-8
from flask import Flask, jsonify, request
import json

app = Flask(__name__)


@app.route('/get', methods=['GET'])
def root():
    return jsonify('''
    {
	"apiVersion": "v1",
	"kind": "ReplicationController",
	"metadata": {
		"name": "iposeidon-system-auth"
	},
	"spec": {
		"replicas": 1,
		"template": {
			"metadata": {
				"labels": {
					"name": "poseidon-system-auth",
					"service": "poseidon-system-auth"
				}
			},
			"spec": {
				"containers": [
					{
						"name": "poseidon-system-auth",
						"image": "192.168.66.59:5000/poseidon-release-1.7.0/poseidon-system-auth:latest",
						"env": [
							{
								"name": "LOG_PATH",
								"value": "/var/log/poseidon"
							},
							{
								"name": "dubbo_port",
								"value": "8811"
							},
							{
								"name": "dubbo_registry",
								"value": "zookeeper://zookeeper.default.svc.cluster.local:2181"
							},
							{
								"name": "dubbo_timeout",
								"value": "5000"
							},
							{
								"name": "management_port",
								"value": "8872"
							},
							{
								"name": "server_port",
								"value": "8871"
							},
							{
								"name": "spring_datasource_password",
								"value": "root123"
							},
							{
								"name": "spring_datasource_url",
								"value": "jdbc:mysql://mysql.default.svc.cluster.local:3306/poseidon_dev?useUnicode=true&useSSL=false&characterEncoding=utf-8"
							},
							{
								"name": "spring_datasource_username",
								"value": "root"
							}
						],
						"ports": [
							{
								"containerPort": 8871
							},
							{
								"containerPort": 8872
							},
							{
								"containerPort": 8811
							}
						]
					}
				]
			}
		}
	}
}
''')


if __name__ == '__main__':
    app.run('0.0.0.0', 60001)
