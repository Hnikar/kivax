"use client";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardTitle, CardDescription } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";

export default function LandingPage() {
	return (
		<div className="min-h-screen flex items-center justify-center bg-slate-50">
			<Card className="w-full max-w-md">
				<CardContent className="space-y-4">
					<CardTitle className="text-lg">Kivax</CardTitle>
					<CardDescription className="text-sm text-slate-600">Enter a code to join or create a new meeting.</CardDescription>

					<input
						aria-label="Meeting code"
						placeholder="Enter meeting code"
						className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-sky-400"
					/>

					<div className="flex gap-2">
						<Button className="flex-1">Join Meeting</Button>
						<Button variant="secondary" className="flex-1">
							Create Meeting
						</Button>
					</div>
				</CardContent>
			</Card>
		</div>
	);
}
