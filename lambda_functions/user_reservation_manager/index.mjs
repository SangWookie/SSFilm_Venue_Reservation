import {SQSClient, SendMessageCommand} from "@aws-sdk/client-sqs"

const sqsClient = new SQSClient({ region: "ap-northeast-2" });

const QUEUE_URL = "https://sqs.ap-northeast-2.amazonaws.com/796973485724/reservation-queue.fifo";

export const handler = async (event) => {
    try {
        await sqsClient.send(new SendMessageCommand({
            QueueUrl: QUEUE_URL,
            MessageBody: JSON.stringify({ message: "Hello from Lambda!" }),
            MessageGroupId: "testGroup"
        }));

        return {
            statusCode: 200,
            body: JSON.stringify({ MessageId: "sent" })
        };
    } catch (error) {
        return {
            statusCode: 500,
            body: JSON.stringify({ error: error.message })
        };
    }
};
