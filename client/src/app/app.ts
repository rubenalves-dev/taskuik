import { AsyncPipe } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Component, effect, inject, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Task, TasksService } from './service/tasks.service';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, AsyncPipe],
    templateUrl: './app.html',
    styleUrl: './app.scss',
})
export class App {
    protected readonly tasksService = inject(TasksService);

    protected readonly title = signal('client');

    protected tasks$!: Observable<Task[]>;

    constructor() {
        effect(() => {
            this.tasks$ = this.tasksService.getTasks();
        });
    }
}
