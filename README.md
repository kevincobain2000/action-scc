![action-scc](https://instachart.coveritup.app/bar?metric=+lines&width=960&height=700&subtitle=kevincobain2000@action-scc&data={"x":["YAML","Go"],"y":[[69,172],[0,2],[2,1],[0,34]],"names":["Code","Comment","Files","Complexity"]}&title=SCC+-+Sloc,+Cloc+and+Code)

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

## Allow permissions

- Make sure you have `README.md` in your repo.
- Make sure `Settings > Actions > General` Workflow permissions - Read and write permissions are checked.

<p align="center">
  <img alt="scc logo" src="https://imgur.com/jysKBFC.png" width="560">
</p>


## Optional options

```yaml
- uses: kevincobain2000/action-scc@v1
  with:
    limit: 7 # Optional. Limit number of languages to show in chart (default is 7)
    filename: README.md # Optional. File to update with chart (default is README.md)
    width: 960 # Optional. Width of chart (default is 960)
    height: 700 # Optional. Height of chart (default is 700)
    instachart_url: # Optional. For self hosting charing service. (default is https://instachart.coveritup.app)
```


## Notes

- This action uses [scc](https://github.com/boyter/scc) to generate the data.
- This action uses [instachart](https://github.com/kevincobain2000/instachart) to generate the chart.
- This action updates `README.md` with url of chart image. Or adds new at the top if first time.
