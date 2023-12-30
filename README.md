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

**Options**

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
- This action updates README url with chart image url in place. And adds new at the top if not present.