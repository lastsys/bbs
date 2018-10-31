var floor = 0x00;
var wall = 0x01;
var redRobot = 0x10;
var blueRobot = 0x80;

function run() {
    var actions = {};

    actions[redRobot] = {
        dx: 1,
        dy: 0
    };
    actions[blueRobot] = {
        dx: -1,
        dy: 0
    };

    return actions;
}
