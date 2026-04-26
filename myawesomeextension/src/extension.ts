import * as vscode from 'vscode';

export function activate(context: vscode.ExtensionContext) {
    console.log('Extension "MyAwesomeExtension" is now active!');

    const disposable = vscode.commands.registerCommand('myawesomeextension.helloWorld', () => {
        vscode.window.showInformationMessage('MyAwesomeExtension: Command executed!');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}
