"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var http_1 = __importDefault(require("http"));
var express_1 = __importDefault(require("express"));
var socket_io_1 = __importDefault(require("socket.io"));
// import onMessage from './socket-events/Message';
// import onDisconnect from './socket-events/Disconnect';
// import onProfileUpdate from './socket-events/ProfileUpdate';
// import onLogin from './socket-events/Login';
var HarmonyDB_1 = require("./HarmonyDB");
var Login_1 = __importDefault(require("./socket-events/Login"));
var Message_1 = __importDefault(require("./socket-events/Message"));
var Register_1 = __importDefault(require("./socket-events/Register"));
var Server = /** @class */ (function () {
    function Server(port) {
        var _this = this;
        this.app = express_1.default();
        this.getSocketServer = function () {
            return _this.SocketServer;
        };
        this.open = function () {
            return new Promise(function (resolve, reject) {
                _this.HTTPServer.listen(_this.port, '0.0.0.0', function () {
                    resolve();
                });
            });
        };
        this.port = port;
        this.HTTPServer = http_1.default.createServer(this.app);
        this.HTTPServer.on('error', this.errorHandler);
        this.SocketServer = socket_io_1.default(this.HTTPServer);
        this.SocketServer.on('connection', function (socket) {
            Login_1.default(socket);
            Message_1.default(socket);
            Register_1.default(socket);
        });
        this.Database = new HarmonyDB_1.HarmonyDB();
        this.app.use(express_1.default.static('public'));
    }
    Server.prototype.errorHandler = function (err) {
        console.log(err.name);
    };
    return Server;
}());
exports.Server = Server;
