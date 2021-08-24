const simpleGit = require('simple-git');
const git = simpleGit();

// chain together tasks to await final result
await git.init().addRemote('origin', '...remote.git');

// or await each step individually
await git.init();
await git.addRemote('origin', '...remote.git')