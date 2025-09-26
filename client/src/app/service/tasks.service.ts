import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    private readonly http = inject(HttpClient);

    private readonly apiUrl = 'http://localhost:8080';

    getTasks(): Observable<Task[]> {
        return this.http
            .get<any>(`${this.apiUrl}/tasks`)
            .pipe(map((data) => data.data.Items || []));
    }

    addTask(title: string) {
        return this.http.post<Task>(
            `${this.apiUrl}/tasks`,
            { Title: title, Status: 0 },
            {
                headers: new HttpHeaders({
                    'Content-Type': 'application/json',
                }),
            }
        );
    }

    updateTaskStatus(task: Task, status: number) {
        return this.http.put<Task>(
            `${this.apiUrl}/tasks/${task.ID}`,
            { Title: task.Title, Status: status },
            {
                headers: new HttpHeaders({
                    'Content-Type': 'application/json',
                }),
            }
        );
    }
}

export interface Task {
    ID: number;
    Title: string;
    Status: number;
}
