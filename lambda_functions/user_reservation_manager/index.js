const { SQSClient, SendMessageCommand } = require("@aws-sdk/client-sqs");

const sqsClient = new SQSClient({ region: "ap-northeast-2" });

const QUEUE_URL =
    "https://sqs.ap-northeast-2.amazonaws.com/796973485724/reservation-queue.fifo";

exports.handler = async (event, context) => {
    var method = event.requestContext.http.method;
    var path = event.rawPath
    var stage = event.requestContext.stage;
    let pathWithoutStage;
    
    if (path.startsWith(`/${stage}/`)) {
        pathWithoutStage = path.substring(stage.length + 1);
    } else {
        pathWithoutStage = path;
    }

    // return {
    //     statusCode: 200,
    //     body: JSON.stringify({
    //         fullpath: path,
    //         path: pathWithoutStage,
    //         stage: stage,
    //         method: method,
    //         // event: event
    //     }),
    // };

    if (pathWithoutStage === "/reservations") {
        if (method === "GET") {
            return {
                statusCode: 200,
                body: JSON.stringify({ message: "this is / GET!" }),
            };
        } else if (method === "POST") {
            return {
                statusCode: 200,
                body: JSON.stringify({ message: "this is / POST!" }),
            };
        }
    }

    try {
        await sqsClient.send(
            new SendMessageCommand({
                QueueUrl: QUEUE_URL,
                MessageBody: JSON.stringify({ message: "Hello from Lambda!" }),
                MessageGroupId: "testGroup",
            })
        );

        return {
            statusCode: 200,
            body: JSON.stringify({ "path": path, "method": method }),
        };
    } catch (error) {
        return {
            statusCode: 500,
            body: JSON.stringify({ error: error.message }),
        };
    }
};
