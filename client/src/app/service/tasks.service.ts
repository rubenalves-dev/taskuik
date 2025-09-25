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
            { title, status: 0 },
            {
                headers: new HttpHeaders({
                    'Content-Type': 'application/json',
                }),
            }
        );
    }
}

export interface Task {
    Id: number;
    Title: string;
    Status: number;
}
