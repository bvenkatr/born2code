const fs = require('fs');
const ini = require('ini');

var config = ini.parse(fs.readFileSync('./config.ini', 'utf-8'));

let argv = require('yargs');

class BaseUtil {
  constructor(name) {
    this.name = name;
    console.log(this.name);

    argv
      .option('lksdfldsf', {
        alias: 'schemasdfdfdf',
        demandOption: true,
        describe: 'Directory path to keep downloaded files',
        type: 'string'
      });
    this.setup_args(argv);
  }

  setup_args() {
  }
}

module.exports = BaseUtil;
// baseUtil = new BaseUtil("venkat");
// console.log(baseUtil.name);
