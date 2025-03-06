const { SQSClient, SendMessageCommand } = require("@aws-sdk/client-sqs");

const QUEUE_URL =
    "https://sqs.ap-northeast-2.amazonaws.com/796973485724/reservation-queue.fifo";

async function sqsMessageSender(content, messageGroup) {
    const sqsClient = new SQSClient({ region: "ap-northeast-2" });

    let result = await sqsClient.send(
        new SendMessageCommand({
            QueueUrl: QUEUE_URL,
            MessageBody: JSON.stringify(content),
            MessageGroupId: messageGroup,
        })
    );
    return result;
}

async function requestReservation(body) {
    try {
        let result = await sqsMessageSender(body, body.venue);
        return result;
    } catch (error) {
        console.error("예약 요청 실패:", error);
        throw error;
    }
}

module.exports = requestReservation;
