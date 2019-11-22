if (CKB.ARGV.length !== 1) {
    throw "Requires only one argument!";
}

var input_index = 0;
var input_coins = 0;
var buffer = new ArrayBuffer(4);
var ret = CKB.CODE.INDEX_OUT_OF_BOUND;

while (true) {
    ret = CKB.raw_load_cell_data(buffer, 0, input_index, CKB.SOURCE.GROUP_INPUT);
    if (ret === CKB.CODE.INDEX_OUT_OF_BOUND) {
        break;
    }
    if (ret !== 4) {
        throw "Invalid input cell!";
    }
    var view = new DataView(buffer);
    input_coins += view.getUint32(0, true);
    input_index += 1;
}

var output_index = 0;
var output_coins = 0;

while (true) {
    ret = CKB.raw_load_cell_data(buffer, 0, output_index, CKB.SOURCE.GROUP_OUTPUT);
    if (ret === CKB.CODE.INDEX_OUT_OF_BOUND) {
        break;
    }
    if (ret !== 4) {
        throw "Invalid output cell!";
    }
    var view = new DataView(buffer);
    output_coins += view.getUint32(0, true);
    output_index += 1;
}

if (input_coins !== output_coins) {
    if (!((input_index === 0) && (output_index === 1))) {
        throw "Invalid token issuing mode!";
    }
    var first_input = CKB.load_input(0, 0, CKB.SOURCE.INPUT);
    if (typeof first_input === "number") {
        throw "Cannot fetch the first input";
    }
    var hex_input = Array.prototype.map.call(
        new Uint8Array(first_input),
        function(x) { return ('00' + x.toString(16)).slice(-2); }).join('');
    if (CKB.ARGV[0] != hex_input) {
        throw "Invalid creation argument!";
    }
}