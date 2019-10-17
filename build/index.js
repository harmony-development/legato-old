"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var http_1 = __importDefault(require("http"));
var express_1 = __importDefault(require("express"));
var socket_io_1 = __importDefault(require("socket.io"));
var chalk_1 = __importDefault(require("chalk"));
var PORT = 4000;
// export these two just in case we need it sometime
exports.app = express_1.default();
exports.httpServer = http_1.default.createServer(exports.app);
exports.socketServer = socket_io_1.default(exports.httpServer);
var clientData = {};
// dummy response
exports.app.get('/', function (req, res) {
    res.json({ bacon: true });
});
exports.socketServer.on('connection', function (socket) {
    socket.on('disconnect', function () {
        exports.socketServer.emit('ClientDisconnectEvent', {
            userid: clientData[socket.id] || socket.id
        });
    });
    socket.on('ClientConnect', function (data) {
        if (data.name) {
            clientData[socket.id] = {
                name: data.name.substr(0, 30)
            };
            exports.socketServer.emit('ClientConnectEvent', {
                userid: clientData[socket.id].name
            });
        }
    });
    socket.on('message', function (message) {
        if (message.message && clientData[socket.id]) {
            exports.socketServer.emit('message', {
                user: clientData[socket.id].name,
                message: message.message
            });
        }
    });
    socket.on('UsernameUpdate', function (name) {
        if (clientData[socket.id]) {
            var oldName = clientData[socket.id].name;
            clientData[socket.id].name = name;
            exports.socketServer.emit('message', {
                user: oldName,
                message: " has changed their name to " + clientData[socket.id].name
            });
        }
    });
});
exports.httpServer.listen(PORT, function () {
    console.log(chalk_1.default.green("Server listening on port " + PORT));
});
