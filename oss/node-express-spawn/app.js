"use strict";

const {spawn} = require('child_process');
const express = require('express');
const app = express();

const {PORT = 3000} = process.env;

app.get('/dir/*', function (req, res) {
  spawnPgm('ls', ['-lh', req.path.split('/dir')[1]], (code, output) => {
    res.set('Content-Type', 'text/plain');
    res.send(output);
  });
});

app.get('/cmd/*', function (req, res) {
  const pgm = req.path.split('/cmd/')[1];
  const args = req.query.args ? req.query.args.split(',') : [];
  //
  spawnPgm(pgm, args, (code, output) => {
    res.set('Content-Type', 'text/plain');
    res.send(output);
  });
});

function spawnPgm(pgm, args, done) {
  const run = spawn(pgm, args);

  let output = '';
  run.stdout.on('data', (data) => {
    output += data + '\n';
  });

  run.stderr.on('data', (data) => {
    output += data + '\n';
  });

  run.on('close', (code) => {
    done(code, output);
  });
}

app.listen(PORT);
