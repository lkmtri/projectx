import { kv } from "@vercel/kv";
import { v4 } from "uuid";
import dayjs from "dayjs";
import type { Job, Status } from "@/types";

const terminalStatuses: Status[] = ["completed", "failed"];

export const getJob = async (uuid: string) => {
  return kv.get<Job>(uuid);
};

export const createJob = async (url: string) => {
  const uuid = v4();
  const job: Job = {
    uuid,
    url,
    status: "in-progress",
    startTime: dayjs().unix(),
    endTime: null,
    error: null,
  };
  await kv.set(uuid, job);
  return job;
};

export const updateJobStatus = async (
  uuid: string,
  status: Status,
  error?: string
) => {
  const job = await getJob(uuid);
  if (!job) {
    return;
  }

  job.status = status;
  job.error = error ?? null;

  if (terminalStatuses.includes(status)) {
    job.endTime = dayjs().unix();
  }

  await kv.set(uuid, job);
  return job;
};
