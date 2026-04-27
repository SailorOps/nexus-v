import * as vscode from 'vscode';

export function activate(context: vscode.ExtensionContext) {
    console.log('Extension "Nexus-Launch-Demo" is now active!');

    const disposable = vscode.commands.registerCommand('openNexusPanel', () => {
        vscode.window.showInformationMessage('Nexus-Launch-Demo: Command executed!');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}
