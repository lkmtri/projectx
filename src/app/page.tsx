import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { cloneGitRepo } from "./actions";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-10">
      <Card className="w-full">
        <form action={cloneGitRepo}>
          <CardHeader>
            <CardTitle>Clone your repository</CardTitle>
            <CardDescription>
              Enter a git repo URL to clone with some side-effects
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              <Label htmlFor="git_url">Git repo URL</Label>
              <Input id="git_url" name="git_url" type="url" required />
            </div>
          </CardContent>
          <CardFooter>
            <Button type="submit">Clone</Button>
          </CardFooter>
        </form>
      </Card>
    </main>
  );
}
