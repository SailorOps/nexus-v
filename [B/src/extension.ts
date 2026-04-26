import * as vscode from 'vscode';

export function activate(context: vscode.ExtensionContext) {
    console.log('Extension "MyAwesomeExtension" is now active!');

    const disposable = vscode.commands.registerCommand('[B.helloWorld', () => {
        vscode.window.showInformationMessage('MyAwesomeExtension: Command executed!');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}
