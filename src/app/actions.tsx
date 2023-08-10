"use server";

import { runJob } from "@/lib/git";
import { createJob, updateJobStatus } from "@/lib/storage";
import { RedirectType } from "next/dist/client/components/redirect";
import { redirect } from "next/navigation";

export async function cloneGitRepo(data: FormData) {
  const url = data.get("git_url") as string;
  const job = await createJob(url);
  runJob(job)
    .then((data) => {
      console.log(data);
      updateJobStatus(job.uuid, "completed", String(data));
    })
    .catch((err) => {
      updateJobStatus(job.uuid, "failed", String(err));
    });

  redirect(`/jobs/${job.uuid}`, RedirectType.push);
}
