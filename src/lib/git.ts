import { Job } from "@/types";
import { exec as _exec } from "child_process";
import git from "isomorphic-git";
import path from "path";
import http from "isomorphic-git/http/node";
import fs from "fs";
import { promisify } from "util";
// import { stderr } from "process";

const writeFile = promisify(fs.writeFile);

const exec = (command: string) =>
  new Promise((res, rej) =>
    _exec(command, (err, stdout, stderr) => {
      if (err) {
        rej(stderr);
        return;
      }
      res(stdout);
    })
  );

export const runJob = async (job: Job) => {
  const dir = path.join(process.cwd(), "jobs", job.uuid);
  const gitDir = path.join(dir, ".git");
  const newFile = path.join(dir, "commit.log");
  const newBranch = `crzy-clone-${job.uuid}`;
  await exec(`git clone --depth 1 ${job.url} ${dir}`);
  await writeFile(newFile, job.uuid, "utf-8");
  await exec(`git --git-dir=${gitDir} checkout -b ${newBranch}`);
  await exec(`git --git-dir=${gitDir} add --all`);
  await exec(`git --git-dir=${gitDir} commit -m 'Commit Job ${job.uuid}'`);
  await exec(`git --git-dir=${gitDir} push origin`);
};

// export const runJob = async (job: Job) => {
//   const dir = path.join(process.cwd(), job.uuid);
//   return await git.clone({
//     fs,
//     http,
//     dir,
//     url: job.url,
//     depth: 1,
//     singleBranch: true,
//   });
// };
