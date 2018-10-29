const WIDTH = 40;
const HEIGHT = 25;

function main() {
    let canvas = {
        element: document.getElementById('canvas'),
        scalingFactor: 1,
        font: loadFont(),
        buffer: (function() {
            let a = new Array(HEIGHT);
            for (let i = 0; i < HEIGHT; i++) {
                a[i] = new Array(WIDTH);
                for (let j = 0; j < WIDTH; j++) {
                    a[i][j] = {char: 32, bg: 0, fg: 1, r: false}
                }
            }
            return a;
        })(),
    };

    initializeCanvas(canvas);
    initializeWebSocket(canvas);
    resizeCanvas(canvas);
}

function loadFont() {
    let font = new Image();
    font.src = '/static/font.png';
    return font
}

function initializeCanvas(canvas) {
    window.addEventListener('resize', function() { resizeCanvas(canvas); });
}

function resizeCanvas(canvas) {
    let charWidth = WIDTH * 8;
    let charHeight = HEIGHT * 8;

    let body = document.getElementById('body');

    let xw = flp2(Math.floor(body.clientWidth / charWidth));
    let yw = flp2(Math.floor(body.clientHeight / charHeight));

    let w = Math.min(xw, yw);

    canvas.element.width = w * charWidth;
    canvas.element.height = w * charHeight;
    canvas.scalingFactor = w;

    renderBuffer(canvas);
}

// Round down to nearest power of 2.
function flp2(x) {
    x = x | (x >> 1);
    x = x | (x >> 2);
    x = x | (x >> 4);
    x = x | (x >> 8);
    x = x | (x >> 16);
    if (x < 1) {
        return 1;
    }
    return x - (x >> 1);
}

function initializeWebSocket(canvas) {
    let origin = location.origin.substring(7);
    let socket = new WebSocket('ws://' + origin + '/ws');
    socket.binaryType = 'arraybuffer';

    socket.onopen = function(event) {
        console.log('Connected to socket.');
        let body = document.getElementById('body');
        body.addEventListener('keydown', function(event) {
            event.preventDefault();
            event.stopPropagation();

            let msg = new Uint8Array(2);
            msg[0] = 0x10; // KeyPress
            msg[1] = event.which;
            socket.send(msg.buffer);
        });
    };

    socket.onclose = function(event) {
        console.log('Closing socket.');
    };

    socket.onerror = function(event) {
        console.log('Error!');
        console.log(event);
    };

    socket.onmessage = function(event) {
        let data = new Uint8Array(event.data);

        switch(data[0]) {
            case 0x01:
                let row, col, i = 1;
                for (row = 0; row < HEIGHT; ++row) {
                    for (col = 0; col < WIDTH; ++col) {
                        canvas.buffer[row][col].char = data[i];
                        canvas.buffer[row][col].fg = data[i+1];
                        canvas.buffer[row][col].bg = data[i+2];
                        canvas.buffer[row][col].r = data[i+3];
                        i += 4;
                    }
                }
                break;
        }
        renderBuffer(canvas);
    };
}

function renderBuffer(canvas) {
    let ctx = canvas.element.getContext('2d');
    ctx.imageSmoothingEnabled = false;
    let pos;
    for (let row = 0; row < HEIGHT; row++) {
        for (let col = 0; col < WIDTH; col++) {
            pos = canvas.buffer[row][col];
            let fg, bg;
            if (canvas.r) {
                fg = pos.bg;
                bg = pos.fg;
            } else {
                fg = pos.fg;
                bg = pos.bg;
            }
            putChar(pos.char, col, row, pos.fg, pos.bg, canvas, ctx);
        }
    }
}

function putChar(char, cx, cy, fg, bg, canvas, ctx) {
    let y = (bg + fg * 16) * 8;
    let x = char * 8;
    ctx.drawImage(canvas.font,
        x, y, 8, 8,
        cx * 8 * canvas.scalingFactor, cy * 8 * canvas.scalingFactor,
        8 * canvas.scalingFactor, 8 * canvas.scalingFactor);
}
