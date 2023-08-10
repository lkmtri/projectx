import { ReactNode } from "react";
import { type Metadata } from "next";
import Link from "next/link";
import dayjs from "dayjs";
import { notFound } from "next/navigation";
import { ArrowLeft, ExternalLink } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { getJob } from "@/lib/storage";
import type { Status } from "@/types";

interface PageProps {
  params: { uuid: string };
}

const statusDisplay: Record<Status, string> = {
  completed: "Completed",
  failed: "Failed",
  "in-progress": "In Progress",
};

interface JobDetail {
  key: string;
  title: ReactNode;
  description: ReactNode;
}

export default async function Job({ params }: PageProps) {
  const job = await getJob(params.uuid);
  if (!job) {
    notFound();
  }

  const details: JobDetail[] = [
    {
      key: "uuid",
      title: "Job ID",
      description: (
        <Link
          className="text-slate-500"
          href={`/jobs/${job.uuid}`}
          target="_blank">
          {job.uuid} <ExternalLink className="inline-block" size={14} />
        </Link>
      ),
    },
    {
      key: "url",
      title: "Git URL",
      description: (
        <a href={job.url} target="_blank" rel="noreferrer">
          <p
            className="inline-block text-slate-500 overflow-hidden text-ellipsis"
            style={{ maxWidth: "calc(100%)" }}>
            {job.url} <ExternalLink className="inline-block" size={14} />
          </p>
        </a>
      ),
    },
    {
      key: "status",
      title: "Status",
      description: statusDisplay[job.status],
    },
    {
      key: "startTime",
      title: "Start time",
      description: dayjs.unix(job.startTime).format("HH:MM MMM-DD-YYYY"),
    },
  ];

  if (job.error) {
    details.push({
      key: "error",
      title: "Error",
      description: <code>{String(job.error)}</code>,
    });
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-10">
      <Card className="w-full">
        <CardHeader>
          <CardTitle>Job Details</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {details.map(({ key, title, description }) => (
              <div key={key}>
                <p className="font-semibold">{title}</p>
                <p>{description}</p>
              </div>
            ))}
          </div>
        </CardContent>
        <CardFooter>
          <Link href="/">
            <Button className="space-x-1" variant="secondary">
              <ArrowLeft size={16} />
              <span>Back</span>
            </Button>
          </Link>
        </CardFooter>
      </Card>
    </main>
  );
}

export async function generateMetadata({
  params,
}: PageProps): Promise<Metadata> {
  const uuid = params.uuid;
  return {
    title: `Job ${uuid}`,
  };
}
