export type Status = "in-progress" | "completed" | "failed";

export interface JobState {
  startTime: number;
  endTime: number | null;
  status: Status;
  error: string | null;
}

export interface Job extends JobState {
  uuid: string;
  url: string;
}
