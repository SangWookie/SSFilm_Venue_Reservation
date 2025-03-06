const requestReservation = require("./requestReservation");

exports.handler = async (event, context) => {
    var method = event.requestContext.http.method;
    var path = event.rawPath;
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

    try {
        if (pathWithoutStage === "/reservations") {
            if (method === "GET") {
                return {
                    statusCode: 200,
                    body: JSON.stringify({ message: "this is / GET!" }),
                };
            } else if (method === "POST") {
                let body = JSON.parse(event.body);
                let result = await requestReservation(body);
                return {
                    statusCode: 200,
                    body: JSON.stringify({ result: result }),
                };
            }
        } else if (pathWithoutStage === "/reservations/unavailable-periods") {
            // return unavailable periods
            if (method === "GET") {
                // query db
            }
        } else if (pathWithoutStage === "/reservations/auto-approval-periods") {
            // return auto approval periods
            if (method === "GET") {
                // query db
            }
        }

        return {
            statusCode: 200,
            body: JSON.stringify({ path: path, method: method }),
        };
    } catch (error) {
        return {
            statusCode: 500,
            body: JSON.stringify({ error: error.message }),
        };
    }
};
