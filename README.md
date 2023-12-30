<p align="center">
  <img alt="scc logo" src="https://imgur.com/QZx9ngs.png" width="160">
</p>

<p align="center">
  Sloc, Code, Complexity - Github action using <a href="https://github.com/boyter/scc" target="_blank">scc</a>
 and <a href="https://github.com/kevincobain2000/instachart" target="_blank">instachart.</a>
</p>

**Quick Setup:** Quickly add `scc` output as chart to your `README.md` with single action.


## Usage

```yaml
- uses: kevincobain2000/action-scc@v1
```

With Options

```yaml
- uses: kevincobain2000/action-scc@v1
  with:
    #default is 7
    limit: 7 # optional. Limit number of languages to show in chart
    #default is README.md
    readme_filename: README.md # optional. File to update with chart
    # default is 960
    width: 960 # optional. Width of chart
    # default is 700
    height: 700 # optional. Height of chart
```

