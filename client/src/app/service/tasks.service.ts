import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    private readonly http = inject(HttpClient);

    private readonly apiUrl = 'http://localhost:8080';

    getTasks() {
        return this.http.get<Task[]>(`${this.apiUrl}/tasks`);
    }
}

export interface Task {
    id: number;
    title: string;
    status: number;
}
