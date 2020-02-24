import click
from click_repl import register_repl


@click.group()
def cli():
    pass


@cli.command()
def hello():
    click.echo("Hello world!")


register_repl(cli)
cli()
